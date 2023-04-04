import axios from 'axios';
import type { CreateDokterRequest } from '../requests/CreateDokterRequest';
import type { DeleteDokterRequest } from '../requests/DeleteDokterRequest';
import type { UpdateDokterRequest } from '../requests/UpdateDokterRequest';
import type { DokterResponses } from '../responses/DokterResponses';

export class DokterService{
    // API
    private readonly url = 'http://api.example.com';

    // CREATE DOKTER
    public async createDokter(request: CreateDokterRequest): Promise<DokterResponses>{
        const response =    await axios.post<DokterResponses>(`${this.url}/users`, request);
        return response.data;
        
    }
    // GET DOKTER 
    public async getAllDokter(): Promise<DokterResponses[]> {
        const response = await axios.get<DokterResponses[]>(`${this.url}/index`);
        return response.data;

    }
    // UPDATE DOKTER
    public async updateDokter(request: UpdateDokterRequest): Promise<DokterResponses> {
        const response = await axios.put<DokterResponses>(`${this.url}/update/${request.id}`, request)
        return response.data;
    }
    // DELETE DOKTER 
    public async deleteDokter(request: DeleteDokterRequest): Promise<void>{
        await axios.delete(`${this.url}/users/${request.id}`)
    }
}