import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  private apiUrl = 'http://localhost:8080';

  constructor(private http: HttpClient) {}

  test(): Observable<any> {
    return this.http.get(`${this.apiUrl}/test`);
  }

  login(data: any): Observable<any> {
    return this.http.post(`${this.apiUrl}/api/auth/login`, data);
  }

  register(data: any) {
    return this.http.post(`${this.apiUrl}/api/auth/register`, data);
  }
  getProfile() {
    const token = localStorage.getItem('token');

    return this.http.get(`${this.apiUrl}/api/profile`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
  }
}
