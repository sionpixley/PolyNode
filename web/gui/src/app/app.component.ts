import { Component, HostListener, OnDestroy, OnInit } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { MatIconRegistry } from '@angular/material/icon';
import { GuiService } from './core/services/gui.service';
import { forkJoin, Observable, Subscription } from 'rxjs';
import { DownloadedComponent } from './core/components/downloaded/downloaded.component';
import { AvailableComponent } from './core/components/available/available.component';
import { NodeVersion } from './core/services/gui.service.models';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, DownloadedComponent, AvailableComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent implements OnInit, OnDestroy {
  // public availableVersions: string[] = [];
  // public downloadedVersions: string[] = ['v18.19.1'];
  public availableVersions: NodeVersion[] = [];
  public downloadedVersions: NodeVersion[] = [{ version: 'v18.19.1', lts: true }];
  public version: string = 'v0.0.0';

  private readonly _sub: Subscription = new Subscription();

  constructor(private readonly _iconRegistry: MatIconRegistry, private readonly _api: GuiService) { }

  public ngOnInit(): void {
    this._iconRegistry.setDefaultFontSetClass('material-symbols-sharp');

    let taskList: Observable<any>[] = [];
    taskList.push(this._api.version());
    taskList.push(this._api.list());
    taskList.push(this._api.search());
    this._sub.add(
      forkJoin(taskList).subscribe(
        {
          next: responses => {
            this.version = responses[0].toString();
            // this.downloadedVersions = responses[1] as string[];
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
