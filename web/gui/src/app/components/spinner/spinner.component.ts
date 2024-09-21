import { Component, Input, OnChanges } from '@angular/core';
import { MatProgressBarModule } from '@angular/material/progress-bar';

@Component({
  selector: 'app-spinner',
  standalone: true,
  imports: [MatProgressBarModule],
  templateUrl: './spinner.component.html',
  styleUrl: './spinner.component.scss'
})
export class SpinnerComponent implements OnChanges {
  @Input({ required: true }) public isLoading: boolean = false;
  @Input() public minAnimationTime: number = 300;

  private _timeoutId: number | null = null;

  public ngOnChanges(): void {
    if(this.isLoading) {
      (document.getElementById('app-spinner-container') as HTMLElement).style.display = 'flex';
      if(!this._timeoutId) {
        this._timeoutId = window.setTimeout(
          () => {
            if(this.isLoading) {
              this._timeoutId = null;
            }
            else {
              this._timeoutId = null;
              (document.getElementById('app-spinner-container') as HTMLElement).style.display = 'none';
            }
          },
          this.minAnimationTime
        );
      }
    }
    else if(!this._timeoutId) {
      (document.getElementById('app-spinner-container') as HTMLElement).style.display = 'none';
    }
  }
}
