import { Component, EventEmitter, Input, OnChanges, Output } from '@angular/core';
import { FormControl, ReactiveFormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
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
    MatInputModule,
    MatFormFieldModule,
    MatButtonModule,
    MatIconModule,
    ReactiveFormsModule
  ],
  templateUrl: './available.component.html',
  styleUrl: './available.component.scss'
})
export class AvailableComponent implements OnChanges {
  @Input({ required: true }) public availableVersions: string[] = [];
  @Input({ required: true }) public ltsVersions: string[] = [];

  @Output() public addButtonEmitter: EventEmitter<string[]> = new EventEmitter();
  @Output() public searchButtonEmitter: EventEmitter<string> = new EventEmitter();

  public allAreSelected = false;
  public readonly columns = ['select', 'version', 'lts'];
  public readonly searchControl: FormControl<string | null> = new FormControl('');

  private _selectedVersions: string[] = [];

  public ngOnChanges(): void {
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
    return this.ltsVersions.includes(version);
  }

  public isSelected(version: string): boolean {
    return this._selectedVersions.includes(version);
  }

  public search(): void {
    this.searchButtonEmitter.emit(this.searchControl.value ?? '');
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
