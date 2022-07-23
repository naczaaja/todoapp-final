import { HTTP_INTERCEPTORS } from "@angular/common/http";
import { RestInterceptor } from "./rest.interceptor";

export const httpInterceptorProviders = [
  { provide: HTTP_INTERCEPTORS, useClass: RestInterceptor, multi: true }
]
