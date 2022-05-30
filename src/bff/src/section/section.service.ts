import { HttpService } from '@nestjs/axios';
import { Injectable } from '@nestjs/common';
import { firstValueFrom } from 'rxjs';
import { Section } from './models/section.models';

@Injectable()
export class SectionService {
  constructor(private httpService: HttpService) {}

  async get(): Promise<Section[]> {
    const res = await firstValueFrom(
      this.httpService.get(`http://api/v1/sections`),
    );
    return res.data.items;
  }

  async find(id: string): Promise<Section> {
    const res = await firstValueFrom(
      this.httpService.get(`http://api/v1/sections/${id}`),
    );
    return res.data.item;
  }

  async create(
    property_id: number,
    type: number,
    name: string,
  ): Promise<Section> {
    const res = await firstValueFrom(
      this.httpService.post(`http://api/v1/sections`, {
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
  ): Promise<Section> {
    const res = await firstValueFrom(
      this.httpService.put(`http://api/v1/sections/${id}`, {
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
      this.httpService.delete(`http://api/v1/sections/${id}`),
    );
    return true;
  }
}
