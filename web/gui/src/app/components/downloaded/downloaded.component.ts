import { Component, Input } from '@angular/core';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatTableModule } from '@angular/material/table';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';

@Component({
  selector: 'app-downloaded',
  standalone: true,
  imports: [
    MatTableModule,
    MatCheckboxModule,
    MatIconModule,
    MatButtonModule
  ],
  templateUrl: './downloaded.component.html',
  styleUrl: './downloaded.component.scss'
})
export class DownloadedComponent {
  @Input({ required: true }) public downloadedVersions: string[] = [];

  public allDownloadedAreSelected: boolean = false;
  public readonly columns: string[] = ['select', 'version', 'isCurrent'];

  private selectedDownloadedVersions: string[] = [];

  public isDownloadedSelected(v: string): boolean {
    return this.selectedDownloadedVersions.includes(v);
  }

  public removeButtonIsDisabled(): boolean {
    return this.selectedDownloadedVersions.length === 0;
  }

  public selectAllDownloadedVersions(): void {
    if(this.allDownloadedAreSelected) {
      this.selectedDownloadedVersions = [];
      this.allDownloadedAreSelected = false;
    }
    else {
      this.selectedDownloadedVersions = Array.from(this.downloadedVersions);
      this.allDownloadedAreSelected = true;
    }
  }

  public selectDownloadedVersion(v: string): void {
    if(this.selectedDownloadedVersions.includes(v)) {
      this.selectedDownloadedVersions = this.selectedDownloadedVersions.filter(ver => ver !== v);
      this.allDownloadedAreSelected = false;
    }
    else {
      this.selectedDownloadedVersions.push(v);
      this.allDownloadedAreSelected = this.downloadedVersions.length === this.selectedDownloadedVersions.length;
    }
  }

  public useButtonIsDisabled(): boolean {
    return this.selectedDownloadedVersions.length !== 1;
  }
}
