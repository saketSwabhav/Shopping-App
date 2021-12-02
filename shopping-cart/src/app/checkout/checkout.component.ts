import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { FormGroup, FormControl, FormBuilder, Validators } from '@angular/forms';
import { Navigation, Router } from '@angular/router';
import { NgbAlert } from '@ng-bootstrap/ng-bootstrap';
import { NgxSpinnerService } from 'ngx-spinner';
import { LoginService } from '../services/login.service';
import { TokenStorageService } from '../services/token-storage.service';

@Component({
  selector: 'app-checkout',
  templateUrl: './checkout.component.html',
  styleUrls: ['./checkout.component.css'],
})
export class CheckoutComponent implements OnInit {
  @Output() closeModal = new EventEmitter();
  checkoutForm!: FormGroup;
  @Input() Product: any
  
  User!: any
  constructor(formBuilder: FormBuilder, private router: Router,
    private tokenStorage: TokenStorageService,private SpinnerServiceck: NgxSpinnerService, private loginService :LoginService) {
    // let nav= this.router.getCurrentNavigation()
    // console.log(nav);
    this.User=this.tokenStorage.getUser();
    

  
   
    // if (nav?.extras && nav.extras.state && nav.extras.state.product) {
    //   this.product = nav.extras.state.product;
    // }
    
    
    this.checkoutForm = formBuilder.group({
     

      email: [null,[Validators.required, Validators.email]],
      address: [null,[Validators.required, Validators.minLength(10)]],
      city: [null,Validators.required],
      state: [null,Validators.required],
      zip: [null,[Validators.required, Validators.minLength(6)]],
    });
  }

  ngOnInit(): void {
    if (this.Product==null) {
      this.router.navigate(['dashboard'])
    }
    console.log("product in co",this.Product);
  }
  postData(){
    alert("Confirm to buy the product")
    console.log(this.checkoutForm.controls);
    this.SpinnerServiceck.show();

    console.log(this.Product.id,this.User.id);
    
    this.loginService.placeOrder(this.Product.id, this.User.id,1,true).subscribe((data)=>
    {
      console.log(data);
      
    })

    setTimeout(() => {
     this.SpinnerServiceck.hide();
     console.log("timer");
     alert("Your Order Placed Successfully")
     this.router.navigate(['/dashboard'])
     this.closeModal.emit();
    }, 3000);
  }
  
}
