import { apiRequest } from './client';

export const profileApi = {
  get: () => 
    apiRequest('/profile'),
    
  update: (profileData) =>
    apiRequest('/profile', {
      method: 'PUT',
      body: JSON.stringify(profileData)
    })
}; 