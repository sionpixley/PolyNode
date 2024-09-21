import { Component, HostListener, OnDestroy, OnInit } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { MatIconRegistry } from '@angular/material/icon';
import { GuiService } from './services/gui.service';
import { forkJoin, Observable, Subscription } from 'rxjs';
import { DownloadedComponent } from './components/downloaded/downloaded.component';
import { AvailableComponent } from './components/available/available.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, DownloadedComponent, AvailableComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent implements OnInit, OnDestroy {
  public availableVersions: string[] = [];
  public currentVersion: string = '';
  public downloadedVersions: string[] = [];
  public version: string = 'v0.0.0';

  private readonly _sub: Subscription = new Subscription();

  constructor(private readonly _iconRegistry: MatIconRegistry, private readonly _api: GuiService) { }

  public ngOnInit(): void {
    this._iconRegistry.setDefaultFontSetClass('material-symbols-sharp');

    let taskList: Observable<any>[] = [];
    taskList.push(this._api.version());
    taskList.push(this._api.list());
    // taskList.push(this._api.search());
    this._sub.add(
      forkJoin(taskList).subscribe(
        {
          next: responses => {
            this.version = responses[0].toString();

            let temp: string[] = responses[1] as string[];
            for(let i = 0; i < temp.length; i += 1) {
              if(temp[i].includes('(current)')) {
                temp[i] = temp[i].replace(' (current)', '');
                this.currentVersion = temp[i];
              }
            }
            this.downloadedVersions = temp;

            // this.availableVersions = responses[2] as string[];
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

  private _cleanup(): void {
    this._sub.unsubscribe();
  }
}
