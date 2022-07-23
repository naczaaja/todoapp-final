import { Injectable } from '@angular/core';
import {
  HttpRequest,
  HttpHandler,
  HttpEvent,
  HttpInterceptor,
  HttpResponse
} from '@angular/common/http';
import { Observable } from 'rxjs';
import { tap } from 'rxjs/operators';
import { environment } from 'src/environments/environment';

@Injectable()
export class RestInterceptor implements HttpInterceptor {

  constructor() { }

  intercept(request: HttpRequest<unknown>, next: HttpHandler): Observable<HttpEvent<unknown>> {
    const url = `${environment.baseURL}${request.url}`
    const urlReq = request.clone({ url })

    console.log("REQ: RestInterceptor");

    return next.handle(urlReq).pipe(
      tap(event => {
        if (event instanceof HttpResponse) {

        }
      })
    );
  }
}
