import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { AccessibilityService } from './core/services/accessibility.service';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet],
  template: '<router-outlet></router-outlet>',
  styles: [':host { display: block; min-height: 100vh; }'],
})
export class AppComponent {
  constructor(_a11y: AccessibilityService) {
    /* load saved display preferences app-wide */
  }
}
