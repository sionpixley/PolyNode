import { Component, Input } from '@angular/core';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatTableModule } from '@angular/material/table';
import { NodeVersion } from '../../services/gui.service.models';

@Component({
  selector: 'app-downloaded',
  standalone: true,
  imports: [MatTableModule, MatCheckboxModule],
  templateUrl: './downloaded.component.html',
  styleUrl: './downloaded.component.scss'
})
export class DownloadedComponent {
  @Input({ required: true }) public downloadedVersions: NodeVersion[] = [];

  public readonly columns: string[] = ['select', 'version', 'lts'];
}
