import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';

export interface Car {
  id: number;
  name: string;
  price: number;
  imageUrl: string;
}

@Injectable({
  providedIn: 'root'
})
export class CarService {
  private cars: Car[] = [
    { id: 1, name: 'Toyota Avanza', price: 300000, imageUrl: 'https://example.com/avanza.jpg' },
    { id: 2, name: 'Honda Jazz', price: 350000, imageUrl: 'https://example.com/jazz.jpg' },
    { id: 3, name: 'Suzuki Ertiga', price: 400000, imageUrl: 'https://example.com/ertiga.jpg' },
  ];

  getCars(): Observable<Car[]> {
    return of(this.cars);
  }

  getCarById(id: number): Observable<Car | undefined> {
    return of(this.cars.find(car => car.id === id));
  }
}
