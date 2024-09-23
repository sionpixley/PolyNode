import { Component, EventEmitter, Input, OnChanges, Output } from '@angular/core';
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

  @Output() public stopLoadingEmitter: EventEmitter<void> = new EventEmitter();

  private _timeoutId: number | null = null;

  public ngOnChanges(): void {
    if(this.isLoading) {
      if(!this._timeoutId) {
        this._timeoutId = window.setTimeout(
          () => {
            if(this.isLoading) {
              this._timeoutId = null;
            }
            else {
              this._timeoutId = null;
              this.stopLoadingEmitter.emit();
            }
          },
          this.minAnimationTime
        );
      }
    }
    else if(!this._timeoutId) {
      this.stopLoadingEmitter.emit();
    }
  }
}
