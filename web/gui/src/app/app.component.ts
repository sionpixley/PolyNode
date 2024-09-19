import { Component, HostListener, OnDestroy, OnInit } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { MatIconRegistry } from '@angular/material/icon';
import { GuiService } from './core/services/gui.service';
import { Subscription } from 'rxjs';
import { DownloadedComponent } from './core/components/downloaded/downloaded.component';
import { AvailableComponent } from './core/components/available/available.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, DownloadedComponent, AvailableComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent implements OnInit, OnDestroy {
  public version: string = 'v0.0.0';

  private readonly _sub: Subscription = new Subscription();

  constructor(private readonly _iconRegistry: MatIconRegistry, private readonly _api: GuiService) { }

  public ngOnInit(): void {
    this._iconRegistry.setDefaultFontSetClass('material-symbols-sharp');

    this._sub.add(
      this._api.version().subscribe(
        {
          next: v => this.version = v,
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
