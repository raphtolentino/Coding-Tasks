import org.epochx.gp.model.*;
import org.epochx.life.*;
import org.epochx.stats.*;

import org.epochx.gp.op.crossover.UniformPointCrossover;
import org.epochx.op.selection.TournamentSelector;
import org.epochx.tools.random.MersenneTwisterFast;

public class Example1
{
    public static void main(String[] args) {
        GPModel model = new EvenParity(4);
        
        // Set parameters.
        model.setPopulationSize(500);
        model.setNoGenerations(100);
        model.setMaxDepth(8);
        
        // Set operators and components.
        model.setCrossover(new UniformPointCrossover(model));
        model.setProgramSelector(new TournamentSelector(model, 7));
        model.setRNG(new MersenneTwisterFast());
        
        Life.get().addGenerationListener(new GenerationAdapter(){
                public void onGenerationEnd() {
                    Stats.get().print(StatField.GEN_NUMBER,
                        StatField.GEN_FITNESS_MIN,
                        StatField.GEN_FITTEST_PROGRAM);
                }
            });
        model.run();
    }
}