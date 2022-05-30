export const SectionTypeEnum = {
  1: "部屋",
  2: "区画",
} as const;

export type sectionTypeEnum = typeof SectionTypeEnum[keyof typeof SectionTypeEnum];
