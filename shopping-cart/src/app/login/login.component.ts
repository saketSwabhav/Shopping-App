
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { LoginService } from '../services/login.service';
import { TokenStorageService } from '../services/token-storage.service';



@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})

export class LoginComponent implements OnInit {

  loginForm!: FormGroup;
  isLoggedIn = false;
  isLoginFailed = false;
  constructor(formBuilder: FormBuilder, private loginService: LoginService,private router: Router,private tokenStorage: TokenStorageService) {

    this.loginForm = formBuilder.group({
      email: ["", [Validators.required, Validators.email]],
      pass: ["", Validators.required],
    });
  }

  ngOnInit(): void {
    if (this.tokenStorage.getToken()!=null) {
      this.router.navigate(['dashboard'])
    }
  }
  postData() {
    console.log(this.loginForm.value.email);
    this.loginService.login(this.loginForm.value.email, this.loginForm.value.pass).subscribe((data) => {
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



