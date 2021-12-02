import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class ProductService {

  private apiUrl ="http://localhost:8080/products"
  constructor(private http: HttpClient) { }
  getProducts(): Observable<any>{
    return this.http.get<any>(this.apiUrl)
  }
  
}
