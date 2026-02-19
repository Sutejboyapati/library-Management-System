import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { LayoutHeaderComponent } from '../layout-header/layout-header.component';
import { LayoutSidebarComponent } from '../layout-sidebar/layout-sidebar.component';

@Component({
  selector: 'app-main-layout',
  standalone: true,
  imports: [RouterOutlet, LayoutHeaderComponent, LayoutSidebarComponent],
  templateUrl: './main-layout.component.html',
  styleUrl: './main-layout.component.css',
})
export class MainLayoutComponent {}
