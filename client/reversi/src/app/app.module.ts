import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms'; // <-- NgModel lives here

import { AppComponent } from './app.component';
import { RoomsComponent } from './rooms/rooms.component';
import { RoomService } from './room.service';
import { AppRoutingModule } from './/app-routing.module';
import { GameComponent } from './game/game.component';


@NgModule({
  declarations: [
    AppComponent,
    RoomsComponent,
    GameComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule
  ],
  providers: [RoomService],
  bootstrap: [AppComponent]
})
export class AppModule { }
