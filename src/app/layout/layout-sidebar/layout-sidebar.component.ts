import { Component } from '@angular/core';
import { RouterLink, RouterLinkActive } from '@angular/router';
import { AuthService } from '../../core/services/auth.service';

@Component({
  selector: 'app-layout-sidebar',
  standalone: true,
  imports: [RouterLink, RouterLinkActive],
  templateUrl: './layout-sidebar.component.html',
  styleUrl: './layout-sidebar.component.css',
})
export class LayoutSidebarComponent {
  constructor(public auth: AuthService) {}
}
