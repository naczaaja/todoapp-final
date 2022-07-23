import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import { AddTodo } from '../interfaces/add-todo';
import { Todo } from '../interfaces/todo';
import { Token } from '../interfaces/token';
import { LocalStorageService } from '../local-starage-service';

@Injectable({
  providedIn: 'root'
})
export class NetworkService {

  constructor(private httpClient: HttpClient, private LocalStorageService: LocalStorageService) { }

  headers = {
    headers: new HttpHeaders()
      .set('Authorization',  `${this.LocalStorageService.get('token')}`)
  }

  userLogin(email:string): Observable<Token> {
    return this.httpClient.post<Token>(`login`,{
      "email" : email
    })
  }

  getUserTodos(): Observable<Todo[]> {
    let headers = {
      headers: new HttpHeaders()
        .set('Authorization',  `${localStorage.getItem('token')}`)
    }
    return this.httpClient.get<Todo[]>(`todos/readtodo`,headers)
  }

  AddUserTodos(addTodo:AddTodo): Observable<Todo> {
    let headers = {
      headers: new HttpHeaders()
        .set('Authorization',  `${localStorage.getItem('token')}`)
    }
    return this.httpClient.post<Todo>(`todos/create`,addTodo,headers)
  }
}
