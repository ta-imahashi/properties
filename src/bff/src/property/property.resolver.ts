import {
  Args,
  ID,
  Mutation,
  Parent,
  Query,
  ResolveField,
  Resolver,
} from '@nestjs/graphql';
import { Section } from 'src/section/models/section.models';
import { Property } from './models/property.models';
import { PropertyService } from './property.service';

@Resolver(() => Property)
export class PropertyResolver {
  constructor(private propertyService: PropertyService) {}

  @Query(() => [Property], { nullable: 'items' })
  getProperty() {
    return this.propertyService.get();
  }

  @Query(() => Property)
  findProperty(@Args('id', { type: () => ID }) id: string) {
    return this.propertyService.find(id);
  }

  @ResolveField('sections', () => [Section])
  withSections(@Parent() parent: Property) {
    return this.propertyService.sections(parent.id);
  }

  @Mutation(() => Property)
  createProperty(
    @Args('property_id') property_id: number,
    @Args('type') type: number,
    @Args('name') name: string,
  ) {
    return this.propertyService.create(property_id, type, name);
  }

  @Mutation(() => Property)
  updateProperty(
    @Args('id', { type: () => ID }) id: number,
    @Args('property_id') property_id: number,
    @Args('type') type: number,
    @Args('name') name: string,
  ) {
    return this.propertyService.update(id, property_id, type, name);
  }

  // mutation updateProperty($id:ID!, $name: String!, $property_id: Float!, $type: Float!){
  //   updateProperty(id:$id, name: $name, property_id:$property_id, type:$type) {
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
  deleteProperty(@Args('id', { type: () => ID }) id: number) {
    return this.propertyService.delete(id);
  }

  // mutation deleteProperty($id:ID!){
  //   deleteProperty(id:$id)
  // }

  // {
  //   "id": 15
  // }
}
