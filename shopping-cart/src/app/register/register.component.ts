import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { LoginService } from '../services/login.service';
import { TokenStorageService } from '../services/token-storage.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {
  registerForm!: FormGroup;
  constructor(formBuilder: FormBuilder,private loginService: LoginService,private router: Router,private tokenStorage: TokenStorageService) {
    this.registerForm = formBuilder.group({
      age: ["", [Validators.required, Validators.min(10), Validators.max(100)]],
      email: ["", [Validators.required, Validators.email]],
      pass: ["", Validators.required],
      conpass: ["", Validators.required],
      name: ["", Validators.required],
      gender: ["", [Validators.required,Validators.nullValidator]],
    });
   }

  ngOnInit(): void {
    if (this.tokenStorage.getToken()!=null) {
      this.router.navigate(['dashboard'])
    }
  }

  postData(){
    console.log(this.registerForm.value.email);
   let fullName = this.registerForm.value.name
    let fn = String(fullName).split(" ") 
    let fName= fn[0]|| ""
    let lname= fn[1]||""
    console.log(fName, lname);
    var valGender=false
    if (this.registerForm.value.age=="Male") {
      valGender=true
    }
    this.loginService.register(this.registerForm.value.email, this.registerForm.value.pass, fName, lname,this.registerForm.value.age,valGender).subscribe((data) => {
      // alert(JSON.parse(JSON.stringify(token.id)))
      this.tokenStorage.saveToken(data.token);
      this.tokenStorage.saveUser(data);
      this.reloadPage();
    },
      err => {

        alert(err.error)
        this.tokenStorage.signOut();
        
      })
    
    
  }
  reloadPage() {
    window.location.reload();
  }
}
