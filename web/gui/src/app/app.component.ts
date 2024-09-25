import { Component, HostListener, OnDestroy, OnInit, signal, WritableSignal } from '@angular/core';
import { MatIconRegistry } from '@angular/material/icon';
import { GuiService } from './services/gui.service';
import { forkJoin, Observable, Subscription } from 'rxjs';
import { DownloadedComponent } from './components/downloaded/downloaded.component';
import { AvailableComponent } from './components/available/available.component';
import { SpinnerComponent } from './components/spinner/spinner.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    DownloadedComponent,
    AvailableComponent,
    SpinnerComponent
  ],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent implements OnInit, OnDestroy {
  public availableVersions: string[] = [];
  public currentVersion = '';
  public downloadedVersions: string[] = [];
  public isLoading = true;
  public ltsVersions: string[] = [];
  public readonly showSpinner: WritableSignal<boolean> = signal(true);
  public version = 'v0.0.0';

  private _addSub: Subscription | null = null;
  private _listSub: Subscription | null = null;
  private _removeSub: Subscription | null = null;
  private readonly _sub = new Subscription();
  private _useSub: Subscription | null = null;

  constructor(private readonly _iconRegistry: MatIconRegistry, private readonly _api: GuiService) { }

  public ngOnInit(): void {
    this._iconRegistry.setDefaultFontSetClass('material-symbols-sharp');

    let taskList: Observable<any>[] = [];
    taskList.push(this._api.version());
    taskList.push(this._api.search());
    this._sub.add(
      forkJoin(taskList).subscribe(
        {
          next: responses => {
            this.reloadDownloadedVersions();
            this.version = responses[0].toString();
            this.availableVersions = responses[1] as string[];
            if(this.availableVersions.length === 14) {
              this.ltsVersions = this.availableVersions.slice(7);
            }
          },
          error: (err: Error) => console.log(err.message)
        }
      )
    );
  }

  public ngOnDestroy(): void {
    this._cleanup();
  }

  @HostListener('window:beforeunload')
  public onRefresh(): void {
    this._cleanup();
  }

  public addButtonClick(versions: string[]): void {
    this.showSpinner.set(true);
    this.isLoading = true;

    let taskList: Observable<boolean>[] = [];
    for(let version of versions) {
      taskList.push(this._api.add(version));
    }

    this._addSub?.unsubscribe();
    this._addSub = forkJoin(taskList).subscribe(
      {
        next: () => this.reloadDownloadedVersions(),
        error: (err: Error) => console.log(err.message)
      }
    );
  }

  public reloadDownloadedVersions(): void {
    this._listSub?.unsubscribe();
    this._listSub = this._api.list().subscribe(
      {
        next: vs => {
          for(let i = 0; i < vs.length; i += 1) {
            if(vs[i].includes('(current)')) {
              vs[i] = vs[i].replace(' (current)', '');
              this.currentVersion = vs[i];
            }
          }
          this.downloadedVersions = vs;
          this.isLoading = false;
        },
        error: (err: Error) => console.log(err.message)
      }
    );
  }

  public removeButtonClick(versions: string[]): void {
    this.showSpinner.set(true);
    this.isLoading = true;

    let taskList: Observable<boolean>[] = [];
    for(let version of versions) {
      taskList.push(this._api.remove(version));
    }

    this._removeSub?.unsubscribe();
    this._removeSub = forkJoin(taskList).subscribe(
      {
        next: () => this.reloadDownloadedVersions(),
        error: (err: Error) => console.log(err.message)
      }
    );
  }

  public stopLoading(): void {
    this.showSpinner.set(false);
  }

  public useButtonClick(selectedVersion: string): void {
    this.showSpinner.set(true);
    this.isLoading = true;
    this._useSub?.unsubscribe();
    this._useSub = this._api.use(selectedVersion).subscribe(
      {
        next: () => this.reloadDownloadedVersions(),
        error: (err: Error) => console.log(err.message)
      }
    );
  }

  private _cleanup(): void {
    this._addSub?.unsubscribe();
    this._addSub = null;

    this._listSub?.unsubscribe();
    this._listSub = null;

    this._removeSub?.unsubscribe();
    this._removeSub = null;

    this._sub.unsubscribe();

    this._useSub?.unsubscribe();
    this._useSub = null;
  }
}
