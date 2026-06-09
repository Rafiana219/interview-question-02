import { Component,OnInit,ChangeDetectorRef } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../../services/auth';
import { CommonModule } from '@angular/common';


@Component({
  selector: 'app-it-02-3',
  standalone: true,
  templateUrl: './it-02-3.html',
  styleUrl: './it-02-3.css',
  imports: [CommonModule],
}) //implements OnInit

export class It023 implements OnInit  {
  
  username = '';

  constructor(
    private router: Router,
    private authService: AuthService,
    private cdr: ChangeDetectorRef,
  ) {}

  ngOnInit(): void {

    const token = localStorage.getItem('token');

    console.log('TOKEN:', token);

    if (!token) {
      this.router.navigate(['/']);
      return;
    }

    this.authService.getProfile().subscribe({
      next: (res: any) => {
        console.log('PROFILE:', res);
        this.username = res.username;
        this.cdr.detectChanges();
        console.log('CURRENT USERNAME:', this.username);
      },
      error: (err) => {
        console.error('PROFILE ERROR:', err);
      },
    });
  }

  logout(): void {
    localStorage.removeItem('token');

    this.router.navigate(['/']);
  }
}
