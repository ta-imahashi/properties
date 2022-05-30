import { Field, ID, ObjectType } from '@nestjs/graphql';
import { Property } from 'src/property/models/property.models';

@ObjectType()
export class Section {
  @Field(() => ID)
  id: string;

  @Field()
  property_id: number;

  @Field(() => Property)
  property: Property;

  @Field()
  type: number;

  @Field()
  name: string;

  @Field()
  created_at: string;

  @Field()
  updated_at: string;
}
