import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { PolyNodeConfig } from './config.service.models';

@Injectable({
  providedIn: 'root'
})
export class ConfigService {
  constructor(private readonly _http: HttpClient) { }

  public loadPolynrc(): Observable<PolyNodeConfig> {
    return this._http.get<PolyNodeConfig>('config/.polynrc');
  }
}
