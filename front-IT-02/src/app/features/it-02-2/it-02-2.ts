import { Component, OnInit ,ChangeDetectorRef } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from '../../services/auth';

@Component({
  selector: 'app-register',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule],
  templateUrl: './it-02-2.html',
  styleUrl: './it-02-2.css',
})
export class It022 implements OnInit {
  registerForm!: FormGroup;

  loading = false;

  errorMessage = '';

  passwordNotMatch = false;

  constructor(
    private fb: FormBuilder,
    private router: Router,
    private authService: AuthService,
    private cdr: ChangeDetectorRef,
  ) {}

  ngOnInit(): void {
    this.registerForm = this.fb.group({
      username: ['', Validators.required],
      password: ['', Validators.required],
      confirmPassword: ['', Validators.required],
    });
  }

 onSubmit(): void {

  if (this.registerForm.invalid) {
    return;
  }

  const password =
    this.registerForm.get('password')?.value;

  const confirmPassword =
    this.registerForm.get('confirmPassword')?.value;

  if (password !== confirmPassword) {

    this.passwordNotMatch = true;
    this.errorMessage = 'รหัสผ่านไม่ตรงกัน';

    return;
  }

  this.passwordNotMatch = false;
  this.errorMessage = '';

  this.loading = true;

  const payload = {
    username: this.registerForm.value.username,
    password: this.registerForm.value.password
  };

  this.authService.register(payload)
    .subscribe({
      next: (res) => {

        console.log('Register Success', res);

        this.loading = false;

        alert('สมัครสมาชิกสำเร็จ');

        this.router.navigate(['/']);
      },

      error: (err) => {

        console.error(err);

        this.loading = false;

        this.errorMessage =
          err.error?.message ||
          'สมัครสมาชิกไม่สำเร็จ';

                
        this.cdr.detectChanges();
      }
    });
}

  goToLogin(): void {
    this.router.navigate(['/']);
  }
}
