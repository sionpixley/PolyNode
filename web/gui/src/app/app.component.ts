import { Component, HostListener, OnDestroy, OnInit, signal, WritableSignal } from '@angular/core';
import { RouterOutlet } from '@angular/router';
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
    RouterOutlet,
    DownloadedComponent,
    AvailableComponent,
    SpinnerComponent
  ],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent implements OnInit, OnDestroy {
  public availableVersions: string[] = ['v18.19.1'];
  public currentVersion: string = '';
  public downloadedVersions: string[] = [];
  public isLoading: boolean = true;
  public version: string = 'v0.0.0';

  private _listSub: Subscription | null = null;
  private _removeSub: Subscription | null = null;
  private readonly _sub: Subscription = new Subscription();
  private _useSub: Subscription | null = null;

  constructor(private readonly _iconRegistry: MatIconRegistry, private readonly _api: GuiService) { }

  public ngOnInit(): void {
    this._iconRegistry.setDefaultFontSetClass('material-symbols-sharp');

    let taskList: Observable<any>[] = [];
    taskList.push(this._api.version());
    // taskList.push(this._api.list());
    // taskList.push(this._api.search());
    this._sub.add(
      forkJoin(taskList).subscribe(
        {
          next: responses => {
            this.version = responses[0].toString();

            // let temp: string[] = responses[1] as string[];
            // for(let i = 0; i < temp.length; i += 1) {
            //   if(temp[i].includes('(current)')) {
            //     temp[i] = temp[i].replace(' (current)', '');
            //     this.currentVersion = temp[i];
            //   }
            // }
            // this.downloadedVersions = temp;

            // this.availableVersions = responses[2] as string[];
          },
          error: (err: Error) => console.log(err.message)
        }
      )
    );

    this.reloadDownloadedVersions();
  }

  public ngOnDestroy(): void {
    this._cleanup();
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

  public useButtonClick(selectedVersion: string): void {
    this.isLoading = true;
    this._useSub?.unsubscribe();
    this._useSub = this._api.use(selectedVersion).subscribe(
      {
        next: () => this.reloadDownloadedVersions(),
        error: (err: Error) => console.log(err.message)
      }
    );
  }

  @HostListener('window:beforeunload')
  public onRefresh(): void {
    this._cleanup();
  }

  private _cleanup(): void {
    this._sub.unsubscribe();
    this._listSub?.unsubscribe();
    this._listSub = null;
    this._removeSub?.unsubscribe();
    this._removeSub = null;
    this._useSub?.unsubscribe();
    this._useSub = null;
  }
}