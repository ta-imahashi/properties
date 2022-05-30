import { HttpService } from '@nestjs/axios';
import { Injectable } from '@nestjs/common';
import { firstValueFrom } from 'rxjs';
import { Section } from 'src/section/models/section.models';
import { Property } from './models/property.models';

@Injectable()
export class PropertyService {
  constructor(private httpService: HttpService) {}

  async get(): Promise<Property[]> {
    const res = await firstValueFrom(
      this.httpService.get(`http://api/v1/properties`),
    );
    return res.data.items;
  }

  async find(id: string): Promise<Property> {
    const res = await firstValueFrom(
      this.httpService.get(`http://api/v1/properties/${id}`),
    );
    return res.data.item;
  }

  async sections(id: string): Promise<Section[]> {
    const res = await firstValueFrom(
      this.httpService.get(`http://api/v1/properties/${id}/sections`),
    );
    return res.data.items;
  }

  async create(
    property_id: number,
    type: number,
    name: string,
  ): Promise<Property> {
    const res = await firstValueFrom(
      this.httpService.post(`http://api/v1/properties`, {
        property_id: property_id,
        type: type,
        name: name,
      }),
    );
    return res.data;
  }

  async update(
    id: number,
    property_id: number,
    type: number,
    name: string,
  ): Promise<Property> {
    const res = await firstValueFrom(
      this.httpService.put(`http://api/v1/properties/${id}`, {
        property_id: property_id,
        type: type,
        name: name,
      }),
    );
    console.log(res);
    return res.data;
  }

  async delete(id: number) {
    await firstValueFrom(
      this.httpService.delete(`http://api/v1/properties/${id}`),
    );
    return true;
  }
}
