import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ProductService } from '../services/product.service';
import {  NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { NgxSpinnerService } from 'ngx-spinner';
import { TokenStorageService } from '../services/token-storage.service';


@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})

export class DashboardComponent implements OnInit {
  userName=""
  User!: string;
  Product : any
  products: any=[]
  
  constructor(private ProductService : ProductService,private router : Router,
    private modalService: NgbModal,private SpinnerService: NgxSpinnerService,private tokenStorage: TokenStorageService) {
   this.SpinnerService.show();
   setTimeout(() => {
    this.SpinnerService.hide();
    console.log("timer");
    
   }, 3000);
    // this.router = _router; 
    this.ProductService.getProducts().subscribe((products) =>
      // console.log(products)
      this.products = products
      )
     
    
  }

  ngOnInit(): void {
    if (this.tokenStorage.getToken()!=null) {
      this.User=this.tokenStorage.getUser();
      console.log(this.User);
      
    } else{
      this.router.navigate(['login'])
    }
    
  }
  headers = ["Item Name", "Item Description", "Item Price"];
  showUpdate(content: any){
    console.log("fd");
    
    // this.modalService.open(content, {ariaLabelledBy: 'modal-basic-title',size: 'xl'}).result.then((result)=>
    // console.log(result)
    // );

    this.openModal(content)
  }
  PostData(content: any,product:any){
    this.Product= product
    // this.router.navigate(['/checkout'],{state: {product :product}});
   this.openModal(content)
    // modalRef.componentInstance.name= product.itemName
    // modalRef.componentInstance.desc= product.itemDesc
    // modalRef.componentInstance.price= product.itemPrice

    
    console.log(product)
  }
  openModal(content: any){
    this.modalService.open(content, {ariaLabelledBy: 'modal-basic-title', size: 'xl'}).result.then((result)=>

    console.log(result)
    );
  }
  
}


