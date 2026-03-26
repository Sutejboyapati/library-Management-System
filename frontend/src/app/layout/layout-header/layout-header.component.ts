import { Component } from '@angular/core';
import { AuthService } from '../../core/services/auth.service';

@Component({
  selector: 'app-layout-header',
  standalone: true,
  templateUrl: './layout-header.component.html',
  styleUrl: './layout-header.component.css',
})
export class LayoutHeaderComponent {
  constructor(public auth: AuthService) {}
}
