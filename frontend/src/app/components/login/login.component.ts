import { Component, OnInit } from '@angular/core';
import { FormControl, Validators } from '@angular/forms';
import { Router, RouterLink } from '@angular/router';
import { isLogin } from 'src/app/local-starage-service';
import { NetworkService } from 'src/app/services/network.service';



@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  emailFormControl = new FormControl('', [Validators.required, Validators.email]);
  isLogin = isLogin

  constructor(private networkService: NetworkService, private router: Router) {

  }

  ngOnInit(): void {
  }

  onClickLogin() {
    let email = this.emailFormControl.value
    this.networkService.userLogin(email).subscribe(result => {
      localStorage.setItem("token", result.token)
      localStorage.setItem("email", result.email)
      this.router.navigate(['/todolist'], {});
    });
  }

}
