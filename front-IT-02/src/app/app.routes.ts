import { Routes } from '@angular/router';

export const routes: Routes = [
  {
    path: '',
    loadComponent: () => import('./features/it-02-1/it-02-1').then((m) => m.It021),
    data: {
      title: 'IT-02-1',
      description: 'Login',
      keyword: 'IT-02-1,login,หน้าแรก',
    },
  },
  {
    path: 'register',
    loadComponent: () => import('./features/it-02-2/it-02-2').then((m) => m.It022),
    data: {
      title: 'IT-02-2',
      description: 'register',
      keyword: 'IT-02-2,register,ลงทะเบียนผู้ใช้ใหม่',
    },
  },
  {
    path: 'dashboard',
    loadComponent: () => import('./features/it-02-3/it-02-3').then((m) => m.It023),
    data: {
      title: 'IT-02-3',
      description: 'dashboard',
      keyword: 'IT-02-3,dashboard,ยินดีต้อนรับ',
    },
  },
  {
    path:'**',
    loadComponent: () => import('./shared/components/not-found/not-found/not-found').then((m)=>m.NotFound)
  },

];
