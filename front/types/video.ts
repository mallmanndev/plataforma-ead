export type Resolution = {
  resolution: string;
  complete_resolution: string;
  url: string;
};

export type Size = {
  low: number;
  high: number;
  unsigned: boolean;
};

export type Video = {
  resolutions: Resolution[];
  id: string;
  type: string;
  status: string;
  size: Size;
  createdAt: string;
  updatedAt: string;
  url: string;
};
