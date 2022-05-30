import { Module } from '@nestjs/common';
import { SectionService } from './section.service';
import { SectionResolver } from './section.resolver';
import { HttpModule } from '@nestjs/axios';
import { PropertyService } from 'src/property/property.service';

@Module({
  imports: [HttpModule],
  providers: [SectionService, SectionResolver, PropertyService],
})
export class SectionModule {}
