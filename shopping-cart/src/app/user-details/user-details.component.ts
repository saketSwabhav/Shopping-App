import { Component, OnInit } from '@angular/core';
import { FormGroup, FormGroupName } from '@angular/forms';

@Component({
  selector: 'app-user-details',
  templateUrl: './user-details.component.html',
  styleUrls: ['./user-details.component.css']
})
export class UserDetailsComponent implements OnInit {

  updateForm!: FormGroup
  constructor() { }

  ngOnInit(): void {
  }

}
