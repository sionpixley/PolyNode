import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class GuiService {
  public guiPort: number = 2334;

  constructor(private readonly _http: HttpClient) { }

  public add(version: string): Observable<boolean> {
    return this._http.post<boolean>(`http://localhost:${this.guiPort}/api/add/${version}`, null);
  }

  public list(): Observable<string[]> {
    return this._http.get<string[]>(`http://localhost:${this.guiPort}/api/list`);
  }

  public remove(version: string): Observable<boolean> {
    return this._http.delete<boolean>(`http://localhost:${this.guiPort}/api/remove/${version}`);
  }

  public search(): Observable<string[]> {
    return this._http.get<string[]>(`http://localhost:${this.guiPort}/api/search`);
  }

  public searchPrefix(prefix: string): Observable<string[]> {
    return this._http.get<string[]>(`http://localhost:${this.guiPort}/api/search/${prefix}`);
  }

  public use(version: string): Observable<boolean> {
    return this._http.patch<boolean>(`http://localhost:${this.guiPort}/api/use/${version}`, null);
  }

  public version(): Observable<string> {
    return this._http.get<string>(`http://localhost:${this.guiPort}/api/version`);
  }
}
