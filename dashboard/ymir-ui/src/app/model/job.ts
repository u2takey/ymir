export class TJob {
  Name: string;
  CurWorkID: string;
  Description: string;
  Script: string;
  ScriptDecode: string;
  Replicas: number;
  NodesSelected: string[];
  Status: string;
  Created: Date;
}

export class TWork {
  JobName: string;
  WorkID: string;
  Replicas: number;
  Result: TResult[];
  Instance: TWorkI[];
}

// TWorkInstance in a single node....
export class TWorkI {
  JobName: string;
  WorkID: string;
  InstanceID: string;
  NodeName: string;
  Created: Date;
  Results: TResult[];
}

// TResult ...
export class TResult {
  Name: string;
  Start: Date;
  End: Date;
  Count: number;
  Sum: number;
  Max: number;
  Avg: number;
  Min: number;
  CodeMap: { [code: number]: number; };// code counter
  Percentile: Bucket[];
  Buckets: Bucket[];
}

// Bucket ...
export class Bucket {
  Count: number;
  Cost: number;
}

// Node ...
export class TNodeList {
  items: TNode[];
}

export class TNode {
  metadata: any;
  checked: boolean;
}

