import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class GuiService {
  private readonly _url = 'http://localhost:2334/api';

  constructor(private readonly _http: HttpClient) { }

  public add(version: string): Observable<boolean> {
    return this._http.post<boolean>(`${this._url}/add/${version}`, null);
  }

  public list(): Observable<string[]> {
    return this._http.get<string[]>(`${this._url}/list`);
  }

  public remove(version: string): Observable<boolean> {
    return this._http.delete<boolean>(`${this._url}/remove/${version}`);
  }

  public search(): Observable<string[]> {
    return this._http.get<string[]>(`${this._url}/search`);
  }

  public searchPrefix(prefix: string): Observable<string[]> {
    return this._http.get<string[]>(`${this._url}/search/${prefix}`);
  }

  public use(version: string): Observable<boolean> {
    return this._http.patch<boolean>(`${this._url}/use/${version}`, null);
  }

  public version(): Observable<string> {
    return this._http.get<string>(`${this._url}/version`);
  }
}
