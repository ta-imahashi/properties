import {
  Args,
  ID,
  Mutation,
  Parent,
  Query,
  ResolveField,
  Resolver,
} from '@nestjs/graphql';
import { Property } from 'src/property/models/property.models';
import { PropertyService } from 'src/property/property.service';
import { Section } from './models/section.models';
import { SectionService } from './section.service';

@Resolver(() => Section)
export class SectionResolver {
  constructor(
    private sectionService: SectionService,
    private propertyService: PropertyService,
  ) {}

  @Query(() => [Section], { nullable: 'items' })
  getSection() {
    return this.sectionService.get();
  }

  @Query(() => Section)
  findSection(@Args('id', { type: () => ID }) id: string) {
    return this.sectionService.find(id);
  }

  @ResolveField('property', () => Property)
  withProperty(@Parent() parent: Section) {
    return this.propertyService.find(parent.property_id + '');
  }

  @Mutation(() => Section)
  createSection(
    @Args('property_id') property_id: number,
    @Args('type') type: number,
    @Args('name') name: string,
  ) {
    return this.sectionService.create(property_id, type, name);
  }

  @Mutation(() => Section)
  updateSection(
    @Args('id', { type: () => ID }) id: number,
    @Args('property_id') property_id: number,
    @Args('type') type: number,
    @Args('name') name: string,
  ) {
    return this.sectionService.update(id, property_id, type, name);
  }

  // mutation updateSection($id:ID!, $name: String!, $property_id: Float!, $type: Float!){
  //   updateSection(id:$id, name: $name, property_id:$property_id, type:$type) {
  //     id
  //   }
  // }

  // {
  //   "id": 4,
  //   "name": "play_update",
  //   "property_id": 1,
  //   "type": 1
  // }

  @Mutation(() => Boolean)
  deleteSection(@Args('id', { type: () => ID }) id: number) {
    return this.sectionService.delete(id);
  }

  // mutation deleteSection($id:ID!){
  //   deleteSection(id:$id)
  // }

  // {
  //   "id": 15
  // }
}
