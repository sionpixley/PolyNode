import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { NodeVersion } from './gui.service.models';

@Injectable({
  providedIn: 'root'
})
export class GuiService {
  constructor(private readonly _http: HttpClient) { }

  public add(version: string): Observable<boolean> {
    return this._http.post<boolean>(`http://localhost:2334/api/add/${version}`, null);
  }

  public list(): Observable<string[]> {
    return this._http.get<string[]>('http://localhost:2334/api/list');
  }

  public remove(version: string): Observable<boolean> {
    return this._http.delete<boolean>(`http://localhost:2334/api/remove/${version}`);
  }

  public search(): Observable<string[]> {
    return this._http.get<string[]>('http://localhost:2334/api/search');
  }

  public searchPrefix(prefix: string): Observable<NodeVersion[]> {
    return this._http.get<NodeVersion[]>('https://nodejs.org/dist/index.json');
  }

  public use(version: string): Observable<boolean> {
    return this._http.patch<boolean>(`http://localhost:2334/api/use/${version}`, null);
  }

  public version(): Observable<string> {
    return this._http.get<string>('http://localhost:2334/api/version');
  }
}
