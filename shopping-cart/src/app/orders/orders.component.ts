import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { LoginService } from '../services/login.service';
import { TokenStorageService } from '../services/token-storage.service';

@Component({
  selector: 'app-orders',
  templateUrl: './orders.component.html',
  styleUrls: ['./orders.component.css']
})
export class OrdersComponent implements OnInit {
  headers = ["Item Name", "Item Description", "Item Price", "Paid","Quantity"];
  User: any;
  orders: any;
  constructor(private tokenStorage: TokenStorageService,private router : Router, private loginService: LoginService) {
   
   }

  ngOnInit(): void {
    if (this.tokenStorage.getToken()!=null) {
      this.User=this.tokenStorage.getUser();
      console.log(this.User);

      this.loginService.getOrders(this.User.id).subscribe((data)=>
    {
      this.orders=data
      console.log(this.orders)
    }
    );
      
    } else{
      this.router.navigate(['login'])
    }
  }

  postData(){
    
  }
}
