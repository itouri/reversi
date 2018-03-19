import { TestBed, inject } from '@angular/core/testing';

import { ReversiService } from './reversi.service';

describe('ReversiService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [ReversiService]
    });
  });

  it('should be created', inject([ReversiService], (service: ReversiService) => {
    expect(service).toBeTruthy();
  }));
});
