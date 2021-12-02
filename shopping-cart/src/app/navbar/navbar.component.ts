import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { Router } from '@angular/router';
import { TokenStorageService } from '../services/token-storage.service';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {
  @Output() updateModal = new EventEmitter();
  constructor(private tokenStorage: TokenStorageService,private router : Router) { }

  ngOnInit(): void {
  }
  logout(){
    this.tokenStorage.signOut();
    console.log("logout");
    
    this.router.navigate(['/login'])
  }
}
