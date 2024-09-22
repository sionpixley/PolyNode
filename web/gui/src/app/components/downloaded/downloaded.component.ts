import { Component, EventEmitter, Input, OnChanges, Output } from '@angular/core';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatTableModule } from '@angular/material/table';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { MatTooltipModule } from '@angular/material/tooltip';

@Component({
  selector: 'app-downloaded',
  standalone: true,
  imports: [
    MatTableModule,
    MatCheckboxModule,
    MatIconModule,
    MatButtonModule,
    MatTooltipModule
  ],
  templateUrl: './downloaded.component.html',
  styleUrl: './downloaded.component.scss'
})
export class DownloadedComponent implements OnChanges {
  @Input({ required: true }) public downloadedVersions: string[] = [];
  @Input({ required: true }) public currentVersion = '';

  @Output() public removeButtonEmitter: EventEmitter<string[]> = new EventEmitter();
  @Output() public useButtonEmitter: EventEmitter<string> = new EventEmitter();

  public allAreSelected = false;
  public readonly columns = ['select', 'version', 'current'];

  private _selectedVersions: string[] = [];

  public ngOnChanges(): void {
    this.downloadedVersions.sort((a, b) => {
      a = a.replace('v', '');
      b = b.replace('v', '');
      let ai = Number.parseInt(a);
      let bi = Number.parseInt(b);
      return bi - ai;
    });
  }

  public isSelected(v: string): boolean {
    return this._selectedVersions.includes(v);
  }

  public removeButtonClick(): void {
    this.removeButtonEmitter.emit(this._selectedVersions);
    this._selectedVersions = [];
    this.allAreSelected = false;
  }

  public removeButtonIsDisabled(): boolean {
    return this._selectedVersions.length === 0;
  }

  public selectAllTooltip(): string {
    return this.allAreSelected ? 'Deselect all' : 'Select all';
  }

  public selectAllVersions(): void {
    if(this.allAreSelected) {
      this._selectedVersions = [];
      this.allAreSelected = false;
    }
    else {
      this._selectedVersions = Array.from(this.downloadedVersions);
      this.allAreSelected = true;
    }
  }

  public selectVersion(v: string): void {
    if(this._selectedVersions.includes(v)) {
      this._selectedVersions = this._selectedVersions.filter(ver => ver !== v);
      this.allAreSelected = false;
    }
    else {
      this._selectedVersions.push(v);
      this.allAreSelected = this.downloadedVersions.length === this._selectedVersions.length;
    }
  }

  public useButtonClick(): void {
    this.useButtonEmitter.emit(this._selectedVersions[0] ?? '');
    this._selectedVersions = [];
  }

  public useButtonIsDisabled(): boolean {
    return this._selectedVersions.length !== 1 || this._selectedVersions[0] === this.currentVersion;
  }

  public useButtonTooltip(): string {
    if(this._selectedVersions.length == 0) {
      return 'Select one row to enable the use button.';
    }
    else if(this._selectedVersions.length > 1) {
      return 'Select only one row to enable the use button.';
    }
    else if(this._selectedVersions[0] === this.currentVersion) {
      return 'Select a Node.js version other than your current one to enable the use button.';
    }
    else {
      return '';
    }
  }
}
