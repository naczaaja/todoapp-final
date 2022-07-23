import { Injectable } from "@angular/core";


@Injectable({
  providedIn: 'root'
})

export class LocalStorageService {
  constructor() { }
  set(key: string, value: string) {
    localStorage.setItem(key, value);
  }

  get(key: string) {
    return localStorage.getItem(key);
  }

  remove(key: string) {
    localStorage.removeItem(key);
  }
}

export function isLogin(): boolean {
  if (localStorage.getItem('token') == null) {
    return false
  } else {
    return true
  }
}
