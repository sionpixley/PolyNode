import { Component, EventEmitter, Input, OnChanges, Output } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatIconModule } from '@angular/material/icon';
import { MatTableModule } from '@angular/material/table';
import { MatTooltipModule } from '@angular/material/tooltip';

@Component({
  selector: 'app-available',
  standalone: true,
  imports: [
    MatTableModule,
    MatCheckboxModule,
    MatButtonModule,
    MatTooltipModule,
    MatIconModule
  ],
  templateUrl: './available.component.html'
})
export class AvailableComponent implements OnChanges {
  @Input({ required: true }) public availableVersions: string[] = [];

  @Output() public addButtonEmitter: EventEmitter<string[]> = new EventEmitter();

  public allAreSelected = false;
  public readonly columns = ['select', 'version', 'lts'];

  private _ltsVersions: string[] = [];
  private _selectedVersions: string[] = [];

  public ngOnChanges(): void {
    if(this.availableVersions.length === 14) {
      this._ltsVersions = this.availableVersions.slice(7);
    }

    this.availableVersions.sort((a, b) => {
      a = a.replace('v', '');
      b = b.replace('v', '');
      let ai = Number.parseInt(a);
      let bi = Number.parseInt(b);
      return bi - ai;
    });
  }

  public addButtonClick(): void {
    this.addButtonEmitter.emit(this._selectedVersions);
    this._selectedVersions = [];
    this.allAreSelected = false;
  }

  public addButtonIsDisabled(): boolean {
    return this._selectedVersions.length === 0;
  }

  public isLts(version: string): boolean {
    return this._ltsVersions.includes(version);
  }

  public isSelected(version: string): boolean {
    return this._selectedVersions.includes(version);
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
      this._selectedVersions = Array.from(this.availableVersions);
      this.allAreSelected = true;
    }
  }

  public selectVersion(version: string): void {
    if(this._selectedVersions.includes(version)) {
      this._selectedVersions = this._selectedVersions.filter(ver => ver !== version);
      this.allAreSelected = false;
    }
    else {
      this._selectedVersions.push(version);
      this.allAreSelected = this._selectedVersions.length === this.availableVersions.length;
    }
  }
}
