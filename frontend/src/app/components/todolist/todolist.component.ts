import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AddTodo } from 'src/app/interfaces/add-todo';
import { Todo } from 'src/app/interfaces/todo';
import { isLogin } from 'src/app/local-starage-service';
import { NetworkService } from 'src/app/services/network.service';


@Component({
  selector: 'app-todolist',
  templateUrl: './todolist.component.html',
  styleUrls: ['./todolist.component.css']
})
export class TodolistComponent implements OnInit {
  isLogin = isLogin
  todoFill?: string;
  todos: Todo[] = []
  email!: string

  constructor(private networkService: NetworkService, private router: Router) { }

  this.email = String(localStorage.getItem('email'))
  ngOnInit(): void {
    if (this.isLogin() == false) {
      this.router.navigate(['/login']);
    }
  }

  ngAfterContentInit(): void {
    //Called after ngOnInit when the component's or directive's content has been initialized.
    //Add 'implements AfterContentInit' to the class.
    this.todos = []
    this.getTodoList()
  }

  getTodoList(): void {
    this.networkService.getUserTodos().subscribe(td => {
      this.todos = td
    });
  }

  addTodo(item: string) {
    let addTodo: AddTodo = {
      message: item
    }
    this.networkService.AddUserTodos(addTodo).subscribe()
    location.reload()
  }

}
