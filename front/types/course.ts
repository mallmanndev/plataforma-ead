export type Item = {
  id: string;
  title: string;
  description: string;
  createdAt: string;
  updatedAt: string;
  videoId: string;
};

export type Section = {
  id: string;
  name: string;
  description: string;
  createdAt: string;
  updatedAt: string;
  itens: Item[];
};

export type Course = {
  id: string;
  name: string;
  description: string;
  visible: boolean;
  createdAt: string;
  updatedAt: string;
  sections: Section[];
};
