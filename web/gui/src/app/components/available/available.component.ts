import { Component, Input, OnInit } from '@angular/core';
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
export class AvailableComponent implements OnInit {
  @Input({ required: true }) public availableVersions: string[] = [];

  public readonly columns: string[] = ['select', 'version', 'lts'];

  private _ltsVersions: string[] = [];

  public ngOnInit(): void {
    if(this.availableVersions.length === 14) {
      this._ltsVersions = this.availableVersions.slice(7);
    }
  }

  public isLts(version: string): boolean {
    return this._ltsVersions.includes(version);
  }
}
