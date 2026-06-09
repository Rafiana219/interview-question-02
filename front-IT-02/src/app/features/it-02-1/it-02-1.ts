import { Component, OnInit,ChangeDetectorRef} from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from '../../services/auth';

@Component({
  selector: 'app-it-02-1',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule],
  templateUrl: './it-02-1.html',
  styleUrl: './it-02-1.css',
})
export class It021 implements OnInit {
  loginForm!: FormGroup;

  loading = false;

  errorMessage = '';

  constructor(
    private fb: FormBuilder,
    private router: Router,
    private authService: AuthService,
    private cdr: ChangeDetectorRef,
  ) {}
  ngOnInit(): void {
    this.loginForm = this.fb.group({
      username: ['', Validators.required],
      password: ['', Validators.required],
    });
  }

  onSubmit(): void {
    if (this.loginForm.invalid) {
      return;
    }

    this.loading = true;
    this.errorMessage = '';

    this.authService.login(this.loginForm.value).subscribe({
      next: (res) => {
        console.log('Login Success', res);

        // เก็บ JWT
        localStorage.setItem('token', res.token);
        console.log(localStorage.getItem('token'));
        

        this.loading = false;

        // ไปหน้า dashboard
        this.router.navigate(['/dashboard']);
      },

      error: (err) => {
        console.error(err);

        this.loading = false;

        this.errorMessage = err.error?.message || 'Login Failed';
        
        this.cdr.detectChanges();
      },
    });
  }



  goToRegister(): void {
    this.router.navigate(['/register']);
  }

  testConnection() {
    this.authService.test().subscribe({
      next: (res) => {
        console.log(res);
        alert(res.message);
      },
      error: (err) => {
        console.error(err);
      },
    });
  }
}
