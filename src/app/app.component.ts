import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'belajar-angular';

  cars = [
    {
      name: 'Toyota Avanza',
      description: 'Toyota Avanza',
      price: 300000,
      imageUrl: '\assets\images\avanza.png' // Sesuaikan dengan path gambar Anda
    },
    {
      name: 'Honda Jazz',
      description: 'Honda Jazz',
      price: 350000,
      imageUrl: '\src\assets\images\jazz.jpg'
    },
    {
      name: 'Suzuki Ertiga',
      description: 'Suzuki Ertiga',
      price: 400000,
      imageUrl: '\assets\images\ertiga.png'
    }
  ];
}
