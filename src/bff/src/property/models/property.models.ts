import { Field, ID, ObjectType } from '@nestjs/graphql';
import { Section } from 'src/section/models/section.models';

@ObjectType()
export class Property {
  @Field(() => ID)
  id: string;

  @Field()
  name: string;

  @Field(() => [Section], { nullable: 'items' })
  sections: Section[];

  @Field()
  created_at: string;

  @Field()
  updated_at: string;
}
