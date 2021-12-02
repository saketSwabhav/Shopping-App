import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type': 'application/json'
  })
}
@Injectable({
  providedIn: 'root'
})
export class LoginService {
  private apiUrl ="http://localhost:8080/"

  constructor(private http: HttpClient) { }
  login(email: string,pass: string): Observable<any>{
    // credentials={
    //   "email": email,
    //   "userPass":pass
    // }
    let body =JSON.stringify({"email":email,"userPass":pass})
    return this.http.post<any>(this.apiUrl+"login", body,httpOptions)
  }

  register(email: string, pass : string,fName: string, lName: string, age: number, gender: boolean): Observable<any>{

    let body = JSON.stringify({"email":email,"userPass":pass,"fName":fName,"lName":lName,"age":age,"isMale":gender})
    return this.http.post<any>(this.apiUrl+"register",body,httpOptions)
  }
  
  placeOrder(productID: string, customerID: string, quantity: number, isPaid:boolean): Observable<any>{
    let body = JSON.stringify({"productID":productID,"customerID":customerID,"quantity":quantity,"isPaid":isPaid})
    return this.http.post<any>(this.apiUrl+"order",body,httpOptions)
  }
  getOrders(customerID: string): Observable<any>{
return this.http.get<any>(this.apiUrl+customerID+"/orders", httpOptions)
  }
}
