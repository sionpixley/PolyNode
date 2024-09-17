import { Routes } from '@angular/router';
import { AvailableComponent } from './core/components/available/available.component';

export const ROUTES: Routes = [
  { path: '', component: AvailableComponent, title: 'Available', pathMatch: 'full' },
  { path: '**', redirectTo: '/' }
];
