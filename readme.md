# goRNG

go implementation of random number generators.

## Pseudo-random number generators (PRNG)

An implementation of the 32-, and 63-bit L'Ecuyer's (1999) pseudo-random number generators.

**Please note** that design of this package follows the work completed by Jochen Voss on his 64bit Mersenne Twister PRNG, see https://github.com/seehuhn/mt19937


## Quantum random number generator (QRNG)

Go client for a 64-bit quantum random number generator service offered by [random.irb.hr](random.irb.hr). *Modified from [https://github.com/salviati/go-qrand](https://github.com/salviati/go-qrand)*. Accesses *"True"* random numbers generated using quantum processes. Requires a username and password.

> We use 'Quantum Random Bit Generator' (QRBG121), which is a fast non-deterministic random bit (number) generator whose randomness relies on intrinsic randomness of the quantum physical process of photonic emission in semiconductors and subsequent detection by photoelectric effect. In this process photons are detected at random, one by one independently of each other. Timing information of detected photons is used to generate random binary digits - bits. The unique feature of this method is that it uses only one photon detector to produce both zeros and ones which results in a very small bias and high immunity to components variation and aging. Furthermore, detection of individual photons is made by a photomultiplier (PMT). Compared to solid state photon detectors the PMT's have drastically superior signal to noise performance and much lower probability of appearing of afterpulses which could be a source of unwanted correlations. Read more at the product page: [http://qrbg.irb.hr/](http://qrbg.irb.hr/). <br> (*from [random.irb.hr](random.irb.hr)*) 


## References

L'Ecuyer, P. 1999. Good parameters and implementations for combined multiple recursive random number generators. Operations Research, 47(1): 159-164.

*see also* Lemieux, C. (2009) Monte Carlo and Quasi-Monte Carlo Sampling. Springer Science. 373pp. [page 63.]
