import { Component, EventEmitter, Input, Output } from '@angular/core';
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
export class DownloadedComponent {
  @Input({ required: true }) public downloadedVersions: string[] = [];
  @Input({ required: true }) public currentVersion: string = '';

  @Output() public removeButtonEmitter: EventEmitter<string[]> = new EventEmitter();
  @Output() public useButtonEmitter: EventEmitter<string> = new EventEmitter();

  public allAreSelected: boolean = false;
  public readonly columns: string[] = ['select', 'version', 'current'];

  private selectedVersions: string[] = [];

  public isSelected(v: string): boolean {
    return this.selectedVersions.includes(v);
  }

  public removeButtonClick(): void {
    this.removeButtonEmitter.emit(this.selectedVersions);
    this.selectedVersions = [];
  }

  public removeButtonIsDisabled(): boolean {
    return this.selectedVersions.length === 0;
  }

  public selectAllTooltip(): string {
    return this.allAreSelected ? 'Deselect all' : 'Select all';
  }

  public selectAllVersions(): void {
    if(this.allAreSelected) {
      this.selectedVersions = [];
      this.allAreSelected = false;
    }
    else {
      this.selectedVersions = Array.from(this.downloadedVersions);
      this.allAreSelected = true;
    }
  }

  public selectVersion(v: string): void {
    if(this.selectedVersions.includes(v)) {
      this.selectedVersions = this.selectedVersions.filter(ver => ver !== v);
      this.allAreSelected = false;
    }
    else {
      this.selectedVersions.push(v);
      this.allAreSelected = this.downloadedVersions.length === this.selectedVersions.length;
    }
  }

  public useButtonClick(): void {
    this.useButtonEmitter.emit(this.selectedVersions[0] ?? '');
    this.selectedVersions = [];
  }

  public useButtonIsDisabled(): boolean {
    return this.selectedVersions.length !== 1 || this.selectedVersions[0] === this.currentVersion;
  }

  public useButtonTooltip(): string {
    if(this.selectedVersions.length == 0) {
      return 'Select one row to enable the use button.';
    }
    else if(this.selectedVersions.length > 1) {
      return 'Select only one row to enable the use button.';
    }
    else if(this.selectedVersions[0] === this.currentVersion) {
      return 'Select a Node.js version other than your current one to enable the use button.';
    }
    else {
      return '';
    }
  }
}
