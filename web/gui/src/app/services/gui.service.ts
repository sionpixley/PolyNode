import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class GuiService {
  public readonly _url = window.location.origin;

  constructor(private readonly _http: HttpClient) { }

  public add(version: string): Observable<boolean> {
    return this._http.post<boolean>(`${this._url}/api/add/${version}`, null);
  }

  public list(): Observable<string[]> {
    return this._http.get<string[]>(`${this._url}/api/list`);
  }

  public remove(version: string): Observable<boolean> {
    return this._http.delete<boolean>(`${this._url}/api/remove/${version}`);
  }

  public search(): Observable<string[]> {
    return this._http.get<string[]>(`${this._url}/api/search`);
  }

  public searchPrefix(prefix: string): Observable<string[]> {
    return this._http.get<string[]>(`${this._url}/api/search/${prefix}`);
  }

  public use(version: string): Observable<boolean> {
    return this._http.patch<boolean>(`${this._url}/api/use/${version}`, null);
  }

  public version(): Observable<string> {
    return this._http.get<string>(`${this._url}/api/version`);
  }
}
